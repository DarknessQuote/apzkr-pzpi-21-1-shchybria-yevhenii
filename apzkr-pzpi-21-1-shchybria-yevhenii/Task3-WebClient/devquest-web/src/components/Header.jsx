import {
	AppBar,
	Box,
	Button,
	ButtonGroup,
	Toolbar,
	Typography,
} from "@mui/material";
import { useAuthContext } from "../context/AuthContext";
import { NavLink, useNavigate } from "react-router-dom";
import { logout } from "../services/authService";
import { useTranslation } from "react-i18next";
import ChangeLanguageMenu from "./ChangeLanguageMenu";
import { dataBackup } from "../services/adminService";

const Header = () => {
	const [auth, setAuth] = useAuthContext();
	const navigate = useNavigate();

	const { t } = useTranslation();

	return (
		<AppBar position="sticky" className="mb-5">
			<Toolbar className="flex items-center gap-10">
				<Box className="grow flex gap-10">
					<NavLink to="/" className="no-underline text-inherit">
						<Typography className="font-bold text-3xl">
							DevQuest
						</Typography>
					</NavLink>
					{auth !== null && auth.role === "Admin" && <AdminActions />}
					{auth !== null && auth.role === "Manager" && (
						<ManagerActions />
					)}
					{auth !== null && auth.role === "Developer" && (
						<DeveloperActions />
					)}
				</Box>
				<ChangeLanguageMenu />
				{auth === null ? (
					<ButtonGroup variant="contained" disableElevation>
						<NavLink to="/auth?mode=login">
							<Button>{t("login")}</Button>
						</NavLink>
						<NavLink to="/auth?mode=register">
							<Button>{t("register")}</Button>
						</NavLink>
					</ButtonGroup>
				) : (
					<Button
						variant="contained"
						onClick={async () => {
							await logout();
							setAuth(null);
							navigate("/");
						}}>
						{t("logout")}
					</Button>
				)}
			</Toolbar>
		</AppBar>
	);
};

const AdminActions = () => {
	const [auth] = useAuthContext();

	const { t } = useTranslation();

	return (
		<ButtonGroup variant="contained" disableElevation>
			<NavLink to="/companies">
				<Button>{t("companies")}</Button>
			</NavLink>
			<Button
				onClick={async () => {
					await dataBackup(auth.token);
				}}>
				{t("backup")}
			</Button>
		</ButtonGroup>
	);
};

const ManagerActions = () => {
	const { t } = useTranslation();

	return (
		<ButtonGroup variant="contained" disableElevation>
			<NavLink to="/projects">
				<Button>{t("projects")}</Button>
			</NavLink>
		</ButtonGroup>
	);
};

const DeveloperActions = () => {
	const { t } = useTranslation();

	return (
		<ButtonGroup variant="contained" disableElevation>
			<NavLink to="/projects">
				<Button>{t("projects")}</Button>
			</NavLink>
		</ButtonGroup>
	);
};

export default Header;
