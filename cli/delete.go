package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

func (r *Operations) DeleteCmd() *cobra.Command {
	var filePath string
	var recursive bool
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a resource",
		Example: `
			bl delete -f ./my-resource.yaml
			# Or using stdin
			cat file.yaml | beamlit delete -f -
		`,
		Run: func(cmd *cobra.Command, args []string) {
			options := map[string]string{
				"environment": environment,
			}
			results, err := getResults(filePath, recursive)
			if err != nil {
				fmt.Printf("error getting results: %v", err)
				os.Exit(1)
			}

			// À ce stade, results contient tous vos documents YAML
			for _, result := range results {
				for _, resource := range resources {
					if resource.Kind == result.Kind {
						name := result.Metadata.(map[string]interface{})["name"].(string)
						resource.DeleteFn(name, options)
					}
				}
			}
		},
	}

	cmd.Flags().BoolVarP(&recursive, "recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	cmd.Flags().StringVarP(&filePath, "filename", "f", "", "containing the resource to delete.")
	err := cmd.MarkFlagRequired("filename")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, resource := range resources {
		subcmd := &cobra.Command{
			Use:     fmt.Sprintf("%s name [flags]", resource.Singular),
			Aliases: []string{resource.Plural, resource.Short},
			Short:   fmt.Sprintf("Delete a %s", resource.Kind),
			Run: func(cmd *cobra.Command, args []string) {
				options := map[string]string{
					"environment": environment,
				}
				if len(args) == 0 {
					fmt.Println("no resource name provided")
					os.Exit(1)
				}
				if len(args) == 1 {
					resource.DeleteFn(args[0], options)
				}
			},
		}
		cmd.AddCommand(subcmd)
	}

	return cmd
}

func (resource Resource) DeleteFn(name string, options map[string]string) {
	ctx := context.Background()
	// Use reflect to call the function
	funcValue := reflect.ValueOf(resource.Delete)
	if funcValue.Kind() != reflect.Func {
		fmt.Println("fn is not a valid function")
		os.Exit(1)
	}
	// Create a slice for the arguments
	fnargs := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(name)} // Add the context and the resource name

	if resource.DeleteParamsType != nil {
		paramsValue := retrieveListParams(resource.DeleteParamsType, options)
		fnargs = append(fnargs, paramsValue)
	}

	// Call the function with the arguments
	results := funcValue.Call(fnargs)

	// Handle the results based on your needs
	if len(results) <= 1 {
		return
	}

	if err, ok := results[1].Interface().(error); ok && err != nil {
		fmt.Println(err)
		return
	}

	// Check if the first result is a pointer to http.Response
	response, ok := results[0].Interface().(*http.Response)
	if !ok {
		fmt.Println("the result is not a pointer to http.Response")
		return
	}
	// Read the content of http.Response.Body
	defer response.Body.Close() // Ensure to close the ReadCloser
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, response.Body); err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode >= 400 {
		ErrorHandler(response.Request, resource.Kind, name, buf.String())
		return
	}

	// Check if the content is an array or an object
	var res interface{}
	if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Resource %s:%s deleted\n", resource.Kind, name)
}
