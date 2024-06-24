import {
	Box,
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	TextField,
} from "@mui/material";
import { useEffect, useRef, useState } from "react";
import { getRoles } from "../services/authService.js";
import { getCompanies } from "../services/companyService.js";
import { useTranslation } from "react-i18next";

const RegisterForm = ({ authenticateUser }) => {
	const [roles, setRoles] = useState([]);
	const [companies, setCompanies] = useState([]);

	const usernameRef = useRef();
	const firstNameRef = useRef();
	const lastNameRef = useRef();
	const passwordRef = useRef();
	const roleRef = useRef();
	const companyRef = useRef();

	const { t } = useTranslation();

	useEffect(() => {
		const getDataForRegister = async () => {
			const roles = await getRoles();
			setRoles(roles);

			const companies = await getCompanies();
			setCompanies(companies);
		};

		getDataForRegister();
	}, []);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const registerData = {
			username: usernameRef.current.value,
			firstName: firstNameRef.current.value,
			lastName: lastNameRef.current.value,
			password: passwordRef.current.value,
			roleID: roleRef.current.value,
			companyID: companyRef.current.value,
		};

		await authenticateUser(registerData);
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
					name="firstName"
					label={t("firstName")}
					variant="outlined"
					inputRef={firstNameRef}
					InputLabelProps={{ shrink: true }}
				/>
				<TextField
					required
					name="lastName"
					label={t("lastName")}
					variant="outlined"
					inputRef={lastNameRef}
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
				<FormControl>
					<InputLabel>{t("role")}</InputLabel>
					<Select
						required
						title="role"
						label={t("role")}
						inputRef={roleRef}>
						<MenuItem selected value="">
							{t("roleSelect")}
						</MenuItem>
						{roles.map((role) => {
							return (
								<MenuItem value={role.id} key={role.id}>
									{t(`${role.title}`)}
								</MenuItem>
							);
						})}
					</Select>
				</FormControl>
				<FormControl>
					<InputLabel>{t("company")}</InputLabel>
					<Select
						required
						title="company"
						label={t("company")}
						inputRef={companyRef}>
						<MenuItem selected value="">
							{t("companySelect")}
						</MenuItem>
						{companies.map((company) => {
							return (
								<MenuItem value={company.id} key={company.id}>
									{company.name}
								</MenuItem>
							);
						})}
					</Select>
				</FormControl>
				<Button type="submit">{t("register")}</Button>
			</Box>
		</form>
	);
};

export default RegisterForm;
