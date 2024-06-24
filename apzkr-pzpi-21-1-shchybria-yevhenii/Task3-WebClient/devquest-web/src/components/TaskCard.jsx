import {
	Button,
	Card,
	CardActions,
	CardContent,
	Divider,
	Typography,
} from "@mui/material";
import { useAuthContext } from "../context/AuthContext";
import { useEffect, useState } from "react";
import { getUser } from "../services/userService";
import { useTranslation } from "react-i18next";

const TaskCard = ({ task, taskActions }) => {
	const [developer, setDeveloper] = useState(null);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			const fetchTaskDeveloper = async () => {
				if (task.developer_id !== null) {
					const fetchedDveloper = await getUser(task.developer_id);
					setDeveloper(fetchedDveloper);
				}
			};

			fetchTaskDeveloper();
		} catch (err) {
			console.error(err);
		}
	}, [task.developer_id]);

	return (
		<Card className="w-4/5 py-3">
			<CardContent>
				<Typography
					variant="h5"
					align={"center"}
					className="mb-3 w-fit transition-all hover:text-gray-800 hover:underline hover:cursor-pointer"
					onClick={() => taskActions.setTaskForInfo(task)}>
					{task.name}
				</Typography>
				<Typography color={"text.secondary"}>
					{t(`${task.category_name}`)}
				</Typography>
				<Divider className="my-3" />
				<Typography>{task.description}</Typography>
				<Typography className="mb-5">
					{t("points")}: {task.points}
				</Typography>
				{task.status_name !== "To do" && developer !== null && (
					<Typography>
						{t("developer")}:{" "}
						{`${developer.last_name} ${developer.first_name}`}
					</Typography>
				)}
			</CardContent>

			{auth !== null && auth.role === "Manager" && (
				<CardActions className="flex justify-center">
					<Button
						variant="contained"
						onClick={() => taskActions.setTaskForEdit(task)}>
						{t("edit")}
					</Button>
					{task.status_name === "To do" && (
						<Button
							variant="contained"
							onClick={() => taskActions.deleteTask(task.id)}>
							{t("delete")}
						</Button>
					)}
				</CardActions>
			)}
			{auth !== null && auth.role === "Developer" && (
				<CardActions className="flex justify-center">
					{task.status_name === "To do" && (
						<Button
							variant="contained"
							onClick={() => taskActions.acceptTask(task.id)}>
							{t("accept")}
						</Button>
					)}
					{task.status_name === "Doing" &&
						task.developer_id === auth.userID && (
							<Button
								variant="contained"
								onClick={() =>
									taskActions.completeTask(task.id)
								}>
								{t("complete")}
							</Button>
						)}
				</CardActions>
			)}
		</Card>
	);
};

export default TaskCard;
