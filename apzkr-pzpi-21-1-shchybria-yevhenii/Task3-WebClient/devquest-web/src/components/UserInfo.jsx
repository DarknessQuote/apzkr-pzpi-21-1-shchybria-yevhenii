import { Card, CardContent, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { getRole } from "../services/userService";
import { getCompany } from "../services/companyService";
import { useTranslation } from "react-i18next";

const UserInfo = ({ userInfo }) => {
	const [userRole, setUserRole] = useState("");
	const [userCompany, setUserCompany] = useState("");

	const { t } = useTranslation();

	useEffect(() => {
		const getRoleAndCompany = async () => {
			const role = await getRole(userInfo.role_id);
			const company = await getCompany(userInfo.company_id);

			setUserRole(role.title);
			setUserCompany(company.name);
		};

		getRoleAndCompany();
	}, [userInfo.company_id, userInfo.role_id]);

	return (
		<Card>
			<CardContent>
				<Typography variant="h5">{`${userInfo.last_name} ${userInfo.first_name}`}</Typography>
				<Typography className="mb-5 text-gray-600">
					{userInfo.username}
				</Typography>
				<Typography>{`${t("role")}: ${t(`${userRole}`)}`}</Typography>
				<Typography>{`${t("company")}: ${userCompany}`}</Typography>
			</CardContent>
		</Card>
	);
};

export default UserInfo;
