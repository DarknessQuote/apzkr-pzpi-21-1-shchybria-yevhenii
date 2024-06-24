import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import { useTranslation } from "react-i18next";
import {
	addAchievement,
	deleteAchievement,
	getProjectAchievements,
	updateAchievement,
} from "../services/achievementService";
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
import AchievementForm from "./AchievementForm";

const AchievementsTab = ({ project }) => {
	const [achievements, setAchievements] = useState([]);
	const [selectedAchievement, setSelectedAchievement] = useState(null);
	const [modalOpen, setModalOpen] = useState(false);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchAchievements = async () => {
					const fetchedAchievements = await getProjectAchievements(
						project.id,
						auth.token
					);
					setAchievements(fetchedAchievements);
				};

				fetchAchievements();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, project.id]);

	const handleOpen = () => setModalOpen(true);
	const handleClose = () => setModalOpen(false);

	const handleDeleteAchievement = async (achievementID) => {
		await deleteAchievement(achievementID, auth.token);
		setAchievements(await getProjectAchievements(project.id, auth.token));
	};

	const sendAchievementData = async (achievementData) => {
		if (achievementData.id.length > 0) {
			await updateAchievement(
				achievementData.id,
				achievementData,
				auth.token
			);
		} else {
			await addAchievement(project.id, achievementData, auth.token);
		}

		setAchievements(await getProjectAchievements(project.id, auth.token));
	};

	return (
		<>
			{auth !== null && auth.role === "Manager" && (
				<Button
					variant="contained"
					className="py-3 mb-8 w-full"
					onClick={() => {
						setSelectedAchievement(null);
						handleOpen();
					}}>
					{t("addAchievement")}
				</Button>
			)}
			<List>
				{achievements.map((achievement, i) => {
					return (
						<Box key={achievement.id}>
							<ListItem className="flex justify-start">
								<ListItemText
									primary={achievement.name}
									secondary={achievement.description}
									className="w-1/2 grow-0"
								/>
								<ListItemText
									primary={`${t("points")}: ${
										achievement.points
									}`}
									className="grow"
								/>
								{auth !== null && auth.role === "Manager" && (
									<ButtonGroup
										variant="contained"
										disableElevation>
										<Button
											onClick={() => {
												setSelectedAchievement(
													achievement
												);
												handleOpen();
											}}>
											{t("edit")}
										</Button>
										<Button
											onClick={() =>
												handleDeleteAchievement(
													achievement.id
												)
											}>
											{t("delete")}
										</Button>
									</ButtonGroup>
								)}
							</ListItem>
							{i < achievements.length - 1 && <Divider />}
						</Box>
					);
				})}
			</List>
			<Dialog open={modalOpen} onClose={handleClose}>
				<DialogContent>
					<AchievementForm
						achievement={selectedAchievement}
						sendAchievementData={sendAchievementData}
						handleClose={handleClose}
					/>
				</DialogContent>
			</Dialog>
		</>
	);
};

export default AchievementsTab;
