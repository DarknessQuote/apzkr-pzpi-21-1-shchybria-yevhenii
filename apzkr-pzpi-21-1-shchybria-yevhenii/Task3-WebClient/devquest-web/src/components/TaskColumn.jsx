import { Stack, Typography } from "@mui/material";
import TaskCard from "./TaskCard";
import { useTranslation } from "react-i18next";

const TaskColumn = ({ tasks, taskStatus, taskActions }) => {
	const { t } = useTranslation()
	
	return (
		<Stack direction={"column"} className="w-1/3 gap-5 items-center">
			<Typography variant="h6" className="mb-5">
				{t(`${taskStatus}`)}
			</Typography>
			{tasks
				.filter((task) => task.status_name === taskStatus)
				.map((task) => {
					return (
						<TaskCard
							task={task}
							taskActions={taskActions}
							key={task.id}
						/>
					);
				})}
		</Stack>
	);
};

export default TaskColumn;
