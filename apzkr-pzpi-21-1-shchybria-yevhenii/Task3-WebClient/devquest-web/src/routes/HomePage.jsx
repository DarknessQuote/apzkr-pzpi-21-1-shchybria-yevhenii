import { useTranslation } from "react-i18next";
import { useAuthContext } from "../context/AuthContext";
import { Box, Button, Divider, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { getUser } from "../services/userService";
import UserInfo from "../components/UserInfo";
import { useNavigate } from "react-router-dom";
import { dataBackup } from "../services/adminService";
import DeveloperAchievements from "../components/DeveloperAchievements";
import DeveloperMeasurements from "../components/DeveloperMeasurements";

const HomePage = () => {
	const [auth] = useAuthContext();
	const [user, setUser] = useState(null);

	const { t } = useTranslation();

	useEffect(() => {
		const getUserInfo = async () => {
			if (auth !== null) {
				const userData = await getUser(auth.userID);
				setUser(userData);
			}
		};

		getUserInfo();
	}, [auth]);

	if (auth == null) {
		return (
			<Typography variant="h5" align="center">
				{t("noAuth")}
			</Typography>
		);
	} else if (user !== null) {
		return (
			<Box>
				<Box className="flex gap-5 items-center">
					<Box className="w-3/5">
						<UserInfo userInfo={user} />
					</Box>
					<Box className="w-2/5 flex flex-col gap-3 justify-stretch items-stretch">
						<Typography
							variant="h5"
							align="center"
							className="mb-3">
							{t("availableActions")}
						</Typography>
						{auth.role === "Admin" && <AdminActions />}
						{auth.role === "Manager" && <ManagerActions />}
						{auth.role === "Developer" && <DeveloperActions />}
					</Box>
				</Box>
				{auth.role === "Developer" && (
					<>
						<Divider className="my-8" />
						<DeveloperAchievements />
						<Divider className="my-8" />
						<DeveloperMeasurements />
					</>
				)}
			</Box>
		);
	}
};

const AdminActions = () => {
	const [auth] = useAuthContext();

	const navigate = useNavigate();

	const { t } = useTranslation();

	return (
		<>
			<Button variant="contained" onClick={() => navigate("/companies")}>
				{t("companies")}
			</Button>
			<Button
				variant="contained"
				onClick={async () => {
					await dataBackup(auth.token);
				}}>
				{t("backup")}
			</Button>
		</>
	);
};

const ManagerActions = () => {
	const navigate = useNavigate();

	const { t } = useTranslation();

	return (
		<Button variant="contained" onClick={() => navigate("/projects")}>
			{t("projects")}
		</Button>
	);
};

const DeveloperActions = () => {
	const navigate = useNavigate();

	const { t } = useTranslation();

	return (
		<Button variant="contained" onClick={() => navigate("/projects")}>
			{t("projects")}
		</Button>
	);
};

export default HomePage;
