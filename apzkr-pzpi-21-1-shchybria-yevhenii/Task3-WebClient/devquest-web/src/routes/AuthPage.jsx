import { useNavigate, useSearchParams } from "react-router-dom";
import { Typography } from "@mui/material";
import LoginForm from "../components/LoginForm.jsx";
import RegisterForm from "../components/RegisterForm.jsx";
import { login, register } from "../services/authService.js";
import { useAuthContext } from "../context/AuthContext.jsx";
import { useTranslation } from "react-i18next";

const AuthPage = () => {
	// eslint-disable-next-line no-unused-vars
	const [_, setAuth] = useAuthContext();
	const navigate = useNavigate();
	const [searchParams] = useSearchParams();
	const authMode = searchParams.get("mode");

	const { t } = useTranslation();

	const authenticateUser = async (authData) => {
		let authResponse;

		if (authMode === "login") {
			authResponse = await login(authData);
		} else if (authMode === "register") {
			authResponse = await register(authData);
		} else {
			console.error("Unsupported auth mode");
			return;
		}

		setAuth({
			token: authResponse.tokens.access_token,
			userID: authResponse.user_id,
			role: authResponse.role,
		});

		navigate("/");
	};

	if (authMode === "login") {
		return <LoginForm authenticateUser={authenticateUser} />;
	} else if (authMode === "register") {
		return <RegisterForm authenticateUser={authenticateUser} />;
	} else {
		return <Typography>{t("invalidAuth")}</Typography>;
	}
};

export default AuthPage;
