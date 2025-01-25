import path from "path";
import { logger } from "../common";
import { AgentBase } from "./base";
import fs from "fs";

export const retrieveWrapperAgent = async (
  dir: string,
  warning: boolean
) => {
  const agents: AgentBase[] = [];
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  
  for (const entry of entries) {
    const fullPath = path.join(dir, entry.name);

    if (entry.isDirectory()) {
      if (entry.name === "node_modules") continue;
      const agentResources = await retrieveWrapperAgent(fullPath, warning)
      agents.push(...agentResources);
    } else if (entry.name.endsWith(".ts") || entry.name.endsWith(".js")) {
      try {
        const module = await import(`${process.cwd()}/${fullPath}`);
        for (const exportedItem of Object.values(module)) {
          const agentBase = (await exportedItem) as AgentBase;
          agents.push(agentBase);
        }
      } catch (error) {
        console.error(error);
        
        if (warning) {
          logger.warn(`Error importing function from ${fullPath}: ${error}`);
        }
      }
    }
  }
  return agents;
};
