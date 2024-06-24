import { useEffect, useState } from "react";
import { getAvailableDevelopers } from "../services/projectService";
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
import { useAuthContext } from "../context/AuthContext";

const AvailableDevelopers = ({ developers, addDeveloper }) => {
	const [availableDevelopers, setAvailableDevelopers] = useState([]);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchAvailableDevelopers = async () => {
					const fetchedDevelopers = await getAvailableDevelopers(
						auth.userID,
						auth.token
					);

					const developersIDs = developers.map(
						(developer) => developer.id
					);
					const available = fetchedDevelopers.filter(
						(developer) =>
							developersIDs.find((id) => id === developer.id) ===
							undefined
					);
					setAvailableDevelopers(available);
				};

				fetchAvailableDevelopers();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, developers]);

	return availableDevelopers.length > 0 ? (
		<List>
			{availableDevelopers.map((developer, i) => {
				return (
					<Box key={developer.id}>
						<ListItem className="flex justify-center gap-10">
							<ListItemText
								primary={`${developer.last_name} ${developer.first_name}`}
								secondary={developer.username}
							/>
							<Button
								variant="contained"
								onClick={() => addDeveloper(developer.id)}>
								{t("add")}
							</Button>
						</ListItem>
						{i < availableDevelopers.length - 1 && <Divider />}
					</Box>
				);
			})}
		</List>
	) : (
		<Typography>{t("noAvailableDevelopers")}</Typography>
	);
};

export default AvailableDevelopers;
