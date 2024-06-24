import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import {
	getDeveloperAchievements,
	getProjectAchievements,
} from "../services/achievementService";
import {
	Box,
	Button,
	Divider,
	List,
	ListItem,
	ListItemText,
	Typography,
} from "@mui/material";
import { useTranslation } from "react-i18next";

const GiveAchievement = ({ projectID, developerID, giveAchievement }) => {
	const [achievements, setAchievements] = useState([]);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchAchievements = async () => {
					const fetchedAchievements = await getProjectAchievements(
						projectID,
						auth.token
					);
					const developerAchievements =
						await getDeveloperAchievements(developerID, auth.token);

					const availableAchievements = fetchedAchievements.filter(
						(achievement) =>
							developerAchievements.find(
								(devAchievement) =>
									devAchievement.id === achievement.id
							) === undefined
					);
					setAchievements(availableAchievements);
				};

				fetchAchievements();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, projectID, developerID]);

	return achievements.length > 0 ? (
		<List>
			{achievements.map((achievement, i) => {
				return (
					<Box key={achievement.id}>
						<ListItem className="flex justify-center gap-10">
							<ListItemText
								primary={achievement.name}
								secondary={achievement.description}
							/>
							<Button
								variant="contained"
								onClick={() => giveAchievement(achievement.id)}>
								{t("giveAchievement")}
							</Button>
						</ListItem>
						{i < achievements.length - 1 && <Divider />}
					</Box>
				);
			})}
		</List>
	) : (
		<Typography>{t("noAvailableAchievements")}</Typography>
	);
};

export default GiveAchievement;
