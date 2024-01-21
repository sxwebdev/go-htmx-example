import "./style.css";
import htmx from "htmx.org";
import Alpine from "alpinejs";
import { WebSocketConnect } from "./dev";

Alpine.start();

// enable hmr
WebSocketConnect();

htmx.config.getCacheBusterParam = true;
