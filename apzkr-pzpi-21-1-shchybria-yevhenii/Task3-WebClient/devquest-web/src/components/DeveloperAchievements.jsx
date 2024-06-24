import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import { getDeveloperAchievements } from "../services/achievementService";
import {
	Box,
	Card,
	CardContent,
	Divider,
	Grid,
	Paper,
	Typography,
} from "@mui/material";
import { getProject } from "../services/projectService";
import { useTranslation } from "react-i18next";

const DeveloperAchievements = () => {
	const [projects, setProjects] = useState([]);
	const [achievements, setAchievements] = useState([]);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchDevAchievments = async () => {
					const fetchedAchievements = await getDeveloperAchievements(
						auth.userID,
						auth.token
					);

					const groupedAchievements = Object.groupBy(
						fetchedAchievements,
						(achievement) => achievement.project_id
					);

					const fetchedProjects = await Promise.all(
						Object.keys(groupedAchievements).map(
							async (projectID) =>
								await getProject(projectID, auth.token)
						)
					);

					setProjects(fetchedProjects);
					setAchievements(fetchedAchievements);
				};

				fetchDevAchievments();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth]);

	return (
		<Box>
			<Typography variant="h5" align={"center"} className="mb-6">
				{t("achievements")}
			</Typography>
			{projects.length > 0 ? (
				projects.map((project) => {
					return (
						<Paper
							elevation={3}
							className="p-3 mb-3"
							key={project.id}>
							<Typography variant="h6">{project.name}</Typography>
							<Divider className="my-3" />
							<Grid container spacing={2}>
								{achievements
									.filter(
										(ach) => ach.project_id === project.id
									)
									.map((ach) => {
										return (
											<Grid item xs={3}>
												<Card>
													<CardContent>
														<Typography className="text-lg">
															{ach.name}
														</Typography>
														<Typography
															color={
																"text.secondary"
															}
															className="text-sm mb-2">
															{ach.description}
														</Typography>
														<Typography>
															{`${t("points")}: ${
																ach.points
															}`}
														</Typography>
													</CardContent>
												</Card>
											</Grid>
										);
									})}
							</Grid>
						</Paper>
					);
				})
			) : (
				<Typography variant="h6" align={"center"}>
					{t("noEarnedAchievements")}
				</Typography>
			)}
		</Box>
	);
};

export default DeveloperAchievements;
