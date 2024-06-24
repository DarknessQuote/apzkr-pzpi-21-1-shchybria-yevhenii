import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import {
	addDeveloperToProject,
	getProjectDevelopers,
	removeDeveloperFromProject,
} from "../services/projectService";
import {
	Box,
	Button,
	ButtonGroup,
	Dialog,
	DialogContent,
	Divider,
	List,
	ListItem,
	ListItemText,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import AvailableDevelopers from "./AvailableDevelopers";
import GiveAchievement from "./GiveAchievement";
import { giveAchievement } from "../services/achievementService";

const DevelopersTab = ({ project }) => {
	const [developers, setDevelopers] = useState([]);
	const [selectedDeveloper, setSelectedDeveloper] = useState(null);
	const [modalOpen, setModalOpen] = useState(false);
	const [modalElement, setModalElement] = useState("developers");

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchDevelopers = async () => {
					const fetchedDevelopers = await getProjectDevelopers(
						project.id,
						auth.token
					);
					setDevelopers(fetchedDevelopers);
				};

				fetchDevelopers();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, project.id]);

	const handleOpen = () => setModalOpen(true);
	const handleClose = () => setModalOpen(false);

	const handleRemoveDeveloper = async (developerID) => {
		await removeDeveloperFromProject(project.id, developerID, auth.token);
		setDevelopers(await getProjectDevelopers(project.id, auth.token));
	};

	const handleAddDeveloper = async (developerID) => {
		await addDeveloperToProject(project.id, developerID, auth.token);
		setDevelopers(await getProjectDevelopers(project.id, auth.token));
		handleClose();
	};

	const handleGiveAchievement = async (achievementID) => {
		await giveAchievement(
			project.id,
			achievementID,
			selectedDeveloper.id,
			auth.token
		);
		setDevelopers(await getProjectDevelopers(project.id, auth.token));

		setModalElement("developers");
		setSelectedDeveloper(null);
		handleClose();
	};

	return (
		<>
			{auth !== null && auth.role === "Manager" && (
				<Button
					variant="contained"
					className="py-3 mb-8 w-full"
					onClick={() => {
						setModalElement("developers");
						setSelectedDeveloper(null);
						handleOpen();
					}}>
					{t("addDeveloper")}
				</Button>
			)}
			<List>
				{developers.map((developer, i) => {
					return (
						<Box key={developer.id}>
							<ListItem className="flex justify-start">
								<ListItemText
									primary={`${developer.last_name} ${developer.first_name}`}
									secondary={developer.username}
									className="w-1/2 grow-0"
								/>
								<ListItemText
									primary={`${t("points")}: ${
										developer.points
									}`}
									className="grow"
								/>
								{auth !== null && auth.role === "Manager" && (
									<ButtonGroup variant="contained">
										<Button
											variant="contained"
											onClick={() => {
												setModalElement("achievements");
												setSelectedDeveloper(developer);
												handleOpen();
											}}>
											{t("giveAchievement")}
										</Button>
										<Button
											variant="contained"
											onClick={() =>
												handleRemoveDeveloper(
													developer.id
												)
											}>
											{t("remove")}
										</Button>
									</ButtonGroup>
								)}
							</ListItem>
							{i < developers.length - 1 && <Divider />}
						</Box>
					);
				})}
			</List>
			<Dialog open={modalOpen} onClose={handleClose}>
				<DialogContent>
					{modalElement === "developers" ? (
						<AvailableDevelopers
							developers={developers}
							addDeveloper={handleAddDeveloper}
						/>
					) : (
						<GiveAchievement
							projectID={project.id}
							developerID={selectedDeveloper.id}
							giveAchievement={handleGiveAchievement}
						/>
					)}
				</DialogContent>
			</Dialog>
		</>
	);
};

export default DevelopersTab;
