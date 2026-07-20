import { createApplication } from "./application.js";
import { loadConfiguration } from "./config.js";

const configuration = loadConfiguration(process.env);
const application = createApplication(configuration);
application.listen();

