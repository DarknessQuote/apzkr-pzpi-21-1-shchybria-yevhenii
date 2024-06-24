import { Box, Button, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

const ProjectForm = ({ project, sendProjectData, handleClose }) => {
	const [id, setID] = useState("");
	const [name, setName] = useState("");
	const [description, setDescription] = useState("");

	const { t } = useTranslation();

	useEffect(() => {
		if (project !== null) {
			setID(project.id);
			setName(project.name);
			setDescription(project.description);
		}
	}, [project]);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const projectData = {
			id: id,
			name: name,
			description: description,
		};

		await sendProjectData(projectData);
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
						name="owner"
						label={t("description")}
						variant="outlined"
						value={description}
						onChange={(e) => setDescription(e.target.value)}
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

export default ProjectForm;
