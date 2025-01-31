import { HumanMessage } from "@langchain/core/messages";
import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { getChatModel, init, logger, wrapFunction } from "../src";

const main = async () => {
  init();
  const getWeather = await wrapFunction(
    (input) => {
      logger.info("Using getWeather");

      if (["sf", "san francisco"].includes(input.location.toLowerCase())) {
        return "It's 60 degrees and foggy.";
      } else {
        return "It's 90 degrees and sunny.";
      }
    },
    {
      name: "get_weather",
      description: "Call to get the current weather.",
      parameters: [
        {
          name: "location",
          description: "Location to get the weather for.",
          type: "string",
        },
      ],
    }
  );

  // const chat = await getChatModel("xai-grok-beta");
  // const chat = await getChatModel("ministral-3b-2410");
  const chat = await getChatModel("gpt-4o-mini");
  // const chat = await getChatModel("cohere-command-r-plus");
  // const chat = await getChatModel("claude-3-5-sonnet");
  // const chat = await getChatModel("deepseek-chat");
  const agent = createReactAgent({
    llm: chat,
    tools: getWeather.tools,
  });
  const result = await agent.invoke({
    messages: [new HumanMessage("What is the weather in San Francisco?")],
  });
  logger.info(result);
};

main();
