import { Box, Button, Divider, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { getUser } from "../services/userService";

const TaskInfo = ({ task, handleClose }) => {
	const [developer, setDeveloper] = useState(null);

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
		<Box>
			<Typography variant="h5" align={"center"} className="mb-5">
				{task.name}
			</Typography>
			<Typography>{task.description}</Typography>
			<Divider className="my-3" />
			<Box className="flex flex-col gap-2">
				<Typography>
					{t("category")}: {t(`${task.category_name}`)}
				</Typography>
				<Typography>
					{t("points")}: {task.points}
				</Typography>
				<Typography>
					{t("expectedTime")}:{" "}
					{new Date(task.expected_time).toLocaleString()}
				</Typography>
				<Typography>
					{t("status")}: {t(`${task.status_name}`)}
				</Typography>
			</Box>
			<Divider className="my-3" />
			<Box className="flex flex-col gap-2 mb-5">
				{task.status_name !== "To do" && developer !== null && (
					<>
						<Typography>
							{t("developer")}:{" "}
							{`${developer.last_name} ${developer.first_name}`}
						</Typography>
						<Typography>
							{t("acceptedTime")}:{" "}
							{new Date(task.accepted_time.Time).toLocaleString()}
						</Typography>
					</>
				)}
				{task.status_name === "Done" && (
					<Typography>
						{t("completedTime")}:{" "}
						{new Date(task.completed_time.Time).toLocaleString()}
					</Typography>
				)}
			</Box>
			<Button
				variant="contained"
				className="w-full"
				onClick={() => handleClose()}>
				{t("close")}
			</Button>
		</Box>
	);
};

export default TaskInfo;
