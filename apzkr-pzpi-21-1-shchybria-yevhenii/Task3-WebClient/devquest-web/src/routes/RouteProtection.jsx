import { Typography } from "@mui/material";
import { useTranslation } from "react-i18next";
import { Outlet } from "react-router-dom";

const RouteProtection = ({ auth, authorizedRoles }) => {
	const { t } = useTranslation();

	if (!auth || !authorizedRoles.includes(auth.role)) {
		return (
			<Typography variant="h6" align={"center"}>
				{t("accessDenied")}
			</Typography>
		);
	}

	return <Outlet />;
};

export default RouteProtection;
