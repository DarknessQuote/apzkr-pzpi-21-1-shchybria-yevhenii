import { Box, Button, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

const AchievementForm = ({ achievement, sendAchievementData, handleClose }) => {
	const [id, setID] = useState("");
	const [name, setName] = useState("");
	const [description, setDescription] = useState("");
	const [points, setPoints] = useState("");

	const { t } = useTranslation();

	useEffect(() => {
		if (achievement !== null) {
			setID(achievement.id);
			setName(achievement.name);
			setDescription(achievement.description);
			setPoints(achievement.points);
		}
	}, [achievement]);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const achievementData = {
			id: id,
			name: name,
			description: description,
			points: Number(points),
		};

		await sendAchievementData(achievementData);
		handleClose();
	};

	return (
		<form onSubmit={handleSubmit}>
			<Box className="flex flex-col items-center gap-5 w-96 py-4">
				<Box className="flex flex-col gap-10 w-4/5">
					<input type="hidden" name="id" value={id} />
					<TextField
						required
						name="name"
						label={t("name")}
						variant="outlined"
						value={name}
						onChange={(e) => setName(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
					<TextField
						required
						name="description"
						label={t("description")}
						variant="outlined"
						value={description}
						onChange={(e) => setDescription(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
					<TextField
						required
						type="number"
						name="points"
						label={t("points")}
						variant="outlined"
						value={points}
						onChange={(e) => setPoints(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
				</Box>
				<Button variant="contained" type="submit">
					{t("save")}
				</Button>
			</Box>
		</form>
	);
};

export default AchievementForm;
