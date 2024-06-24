import { Box, Divider, Tab, Tabs, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getProject } from "../services/projectService";
import { useAuthContext } from "../context/AuthContext";
import { useTranslation } from "react-i18next";
import TasksTab from "../components/TasksTab";
import DevelopersTab from "../components/DevelopersTab";
import AchievementsTab from "../components/AchievementTab";

const ProjectPage = () => {
	const [project, setProject] = useState(null);
	const [tabIndex, setTabIndex] = useState(0);

	const [auth] = useAuthContext();

	const { id } = useParams();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchProjectInfo = async () => {
					const fetchedProject = await getProject(id, auth.token);
					setProject(fetchedProject);
				};

				fetchProjectInfo();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, id]);

	return (
		<>
			<Box>
				<Typography variant="h5" className="mb-2">
					{project?.name}
				</Typography>
				<Typography>{project?.description}</Typography>
				<Divider className="my-4" />

				<Tabs
					value={tabIndex}
					onChange={(_, i) => setTabIndex(i)}
					className="mb-5">
					<Tab label={t("tasks")} />
					<Tab label={t("developers")} />
					<Tab label={t("achievements")} />
				</Tabs>
				{tabIndex === 0 && project !== null && (
					<TasksTab project={project} />
				)}
				{tabIndex === 1 && project !== null && (
					<DevelopersTab project={project} />
				)}
				{tabIndex === 2 && project !== null && (
					<AchievementsTab project={project} />
				)}
			</Box>
		</>
	);
};

export default ProjectPage;
