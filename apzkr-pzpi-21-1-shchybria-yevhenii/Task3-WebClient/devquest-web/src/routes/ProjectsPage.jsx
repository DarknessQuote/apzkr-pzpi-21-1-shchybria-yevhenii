import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import {
	addProject,
	deleteProject,
	getDeveloperProjects,
	getManagerProjects,
	updateProject,
} from "../services/projectService";
import ProjectCard from "../components/ProjectCard";
import { Button, Dialog, DialogContent, Grid } from "@mui/material";
import { useTranslation } from "react-i18next";
import ProjectForm from "../components/ProjectForm";

const ProjectsPage = () => {
	const [projects, setProjects] = useState([]);
	const [selectedProject, setSelectedProject] = useState(null);
	const [modalOpen, setModalOpen] = useState(false);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		const fetchProjects = async () => {
			try {
				if (auth !== null) {
					const fetchedProjects =
						auth.role === "Manager"
							? await getManagerProjects(auth.userID, auth.token)
							: await getDeveloperProjects(
									auth.userID,
									auth.token
							  );
					setProjects(fetchedProjects);
				}
			} catch (err) {
				console.error(err);
			}
		};

		fetchProjects();
	}, [auth, auth?.userID, auth?.token]);

	const handleOpen = () => setModalOpen(true);
	const handleClose = () => setModalOpen(false);

	const sendProjectData = async (projectData) => {
		if (projectData.id.length > 0) {
			await updateProject(projectData.id, projectData, auth.token);
		} else {
			await addProject(auth.userID, projectData, auth.token);
		}

		setProjects(
			auth.role === "Manager"
				? await getManagerProjects(auth.userID, auth.token)
				: await getDeveloperProjects(auth.userID, auth.token)
		);
	};

	const handleProjectDelete = async (projectID) => {
		await deleteProject(projectID, auth.token);
		setProjects(
			auth.role === "Manager"
				? await getManagerProjects(auth.userID, auth.token)
				: await getDeveloperProjects(auth.userID, auth.token)
		);
	};

	return (
		<>
			<Grid container spacing={2}>
				{projects.map((project) => {
					return (
						<Grid item xs={4} key={project.id}>
							<ProjectCard
								project={project}
								handleDelete={handleProjectDelete}
								handleEdit={() => {
									setSelectedProject(project);
									handleOpen();
								}}
								key={project.id}
							/>
						</Grid>
					);
				})}
				{auth !== null && auth.role === "Manager" && (
					<Grid item xs={12}>
						<Button
							variant="contained"
							className="py-3 w-full"
							onClick={() => {
								setSelectedProject(null);
								handleOpen();
							}}>
							{t("addProject")}
						</Button>
					</Grid>
				)}
			</Grid>
			<Dialog open={modalOpen} onClose={handleClose}>
				<DialogContent>
					<ProjectForm
						project={selectedProject}
						sendProjectData={sendProjectData}
						handleClose={handleClose}
					/>
				</DialogContent>
			</Dialog>
		</>
	);
};

export default ProjectsPage;
