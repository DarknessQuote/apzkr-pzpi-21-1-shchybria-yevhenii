import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import {
	CssBaseline,
	StyledEngineProvider,
	ThemeProvider,
	createTheme,
} from "@mui/material";
import { AuthProvider } from "./context/AuthContext";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import "./adapting/i18n";
import { BrowserRouter } from "react-router-dom";

const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement);

const theme = createTheme({
	components: {
		MuiPopover: {
			defaultProps: {
				container: rootElement,
			},
		},
		MuiPopper: {
			defaultProps: {
				container: rootElement,
			},
		},
		MuiDialog: {
			defaultProps: {
				container: rootElement,
			},
		},
	},
});

root.render(
	<StyledEngineProvider injectFirst>
		<ThemeProvider theme={theme}>
			<LocalizationProvider dateAdapter={AdapterDayjs}>
				<AuthProvider>
					<BrowserRouter>
						<CssBaseline />
						<App />
					</BrowserRouter>
				</AuthProvider>
			</LocalizationProvider>
		</ThemeProvider>
	</StyledEngineProvider>
);
