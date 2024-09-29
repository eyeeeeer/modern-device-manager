// Import dependencies
import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import { FluentProvider, webDarkTheme } from "@fluentui/react-components";

const container = document.getElementById("appRoot");

const root = createRoot(container!);

root.render(
    <React.StrictMode>
        <FluentProvider theme={webDarkTheme}>
            <App/>
        </FluentProvider>
    </React.StrictMode>
);