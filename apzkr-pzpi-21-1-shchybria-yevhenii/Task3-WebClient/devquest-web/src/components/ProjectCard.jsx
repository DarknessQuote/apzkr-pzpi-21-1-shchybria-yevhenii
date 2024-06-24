import { useEffect, useState } from "react";
import { getUser } from "../services/userService";
import {
	Button,
	Card,
	CardActions,
	CardContent,
	Typography,
} from "@mui/material";
import { useAuthContext } from "../context/AuthContext";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

const ProjectCard = ({ project, handleEdit, handleDelete }) => {
	const [manager, setManager] = useState(null);

	const [auth] = useAuthContext();

	const navigate = useNavigate();

	const { t } = useTranslation();

	useEffect(() => {
		const fetchManager = async () => {
			try {
				const fetchedManager = await getUser(project.manager_id);
				setManager(fetchedManager);
			} catch (err) {
				console.error(err);
			}
		};

		fetchManager();
	}, [project.manager_id]);

	return (
		<Card>
			<CardContent>
				<Typography
					variant="h5"
					className="mb-1 w-fit transition-all hover:text-gray-800 hover:underline hover:cursor-pointer"
					onClick={() => navigate(`/project/${project.id}`)}>
					{project.name}
				</Typography>
				<Typography className="mb-5">{project.description}</Typography>
				<Typography>
					{t("manager")}:{" "}
					{manager !== null
						? `${manager?.last_name} ${manager?.first_name}`
						: ""}
				</Typography>
			</CardContent>
			{auth !== null && auth.role === "Manager" && (
				<CardActions className="flex justify-center">
					<Button
						variant="contained"
						onClick={() => handleEdit(project)}>
						{t("edit")}
					</Button>
					<Button
						variant="contained"
						onClick={async () => await handleDelete(project.id)}>
						{t("delete")}
					</Button>
				</CardActions>
			)}
		</Card>
	);
};

export default ProjectCard;
