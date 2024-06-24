import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import { useTranslation } from "react-i18next";
import { getMeasurementsForDeveloper } from "../services/measurementService";
import { Box, Card, CardContent, Grid, Typography } from "@mui/material";

const DeveloperMeasurements = () => {
	const [measurements, setMeasurements] = useState([]);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchMeasurements = async () => {
					const fetchedMeasurements =
						await getMeasurementsForDeveloper(auth.userID);
					setMeasurements(fetchedMeasurements);
				};

				fetchMeasurements();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth]);

	return (
		<Box>
			<Typography variant="h5" align={"center"} className="mb-6">
				{t("measurements")}
			</Typography>
			<Grid container spacing={2}>
				{measurements.length > 0 ? (
					measurements.map((measure) => {
						return (
							<Grid item xs={4}>
								<Card>
									<CardContent>
										<Typography>{`${t("measureType")}: ${t(
											measure.type_name
										)}`}</Typography>
										<Typography className="mb-3">{`${t(
											"measureTime"
										)}: ${new Date(
											measure.measured_at
										).toLocaleString()}`}</Typography>
										<Typography>{`${t("measureValue")}: ${
											measure.value
										}`}</Typography>
										<Typography>
											{t(measure.message)}
										</Typography>
									</CardContent>
								</Card>
							</Grid>
						);
					})
				) : (
					<Grid item xs={12}>
						<Typography variant="h6" align={"center"}>
							{t("noMeasurements")}
						</Typography>
					</Grid>
				)}
			</Grid>
		</Box>
	);
};

export default DeveloperMeasurements;
