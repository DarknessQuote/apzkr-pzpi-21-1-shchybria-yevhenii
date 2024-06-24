import { Box, Button, TextField } from "@mui/material";
import { useRef } from "react";
import { useTranslation } from "react-i18next";

const LoginForm = ({ authenticateUser }) => {
	const usernameRef = useRef();
	const passwordRef = useRef();

	const { t } = useTranslation();

	const handleSubmit = async (e) => {
		e.preventDefault();

		const loginData = {
			username: usernameRef.current.value,
			password: passwordRef.current.value,
		};

		await authenticateUser(loginData);
	};

	return (
		<form onSubmit={handleSubmit}>
			<Box className="flex flex-col gap-3">
				<TextField
					required
					name="username"
					label={t("username")}
					variant="outlined"
					inputRef={usernameRef}
					InputLabelProps={{ shrink: true }}
				/>
				<TextField
					required
					name="password"
					type="password"
					label={t("password")}
					variant="outlined"
					inputRef={passwordRef}
					InputLabelProps={{ shrink: true }}
				/>
				<Button type="submit">{t("login")}</Button>
			</Box>
		</form>
	);
};

export default LoginForm;
