import {
	Box,
	Button,
	Dialog,
	DialogContent,
	Divider,
	Stack,
} from "@mui/material";
import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import {
	acceptTask,
	addTask,
	completeTask,
	deleteTask,
	getProjectTasks,
	updateTask,
} from "../services/taskService";
import TaskColumn from "../components/TaskColumn";
import TaskForm from "../components/TaskForm";
import TaskInfo from "../components/TaskInfo";
import { useTranslation } from "react-i18next";

const TasksTab = ({ project }) => {
	const [tasks, setTasks] = useState([]);
	const [selectedTask, setSelectedTask] = useState(null);
	const [modalOpen, setModalOpen] = useState(false);
	const [taskElement, setTaskElement] = useState("form");

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		try {
			if (auth !== null) {
				const fetchProjectTasks = async () => {
					const fetchedTasks = await getProjectTasks(
						project.id,
						auth.token
					);
					setTasks(fetchedTasks);
				};

				fetchProjectTasks();
			}
		} catch (err) {
			console.error(err);
		}
	}, [auth, project.id]);

	const handleOpen = () => setModalOpen(true);
	const handleClose = () => setModalOpen(false);

	const setTaskForInfo = (task) => {
		setSelectedTask(task);
		setTaskElement("info");
		handleOpen();
	};
	const setTaskForEdit = (task) => {
		setSelectedTask(task);
		setTaskElement("form");
		handleOpen();
	};
	const handleDeleteTask = async (taskID) => {
		await deleteTask(taskID, auth.token);
		setTasks(await getProjectTasks(project.id, auth.token));
	};
	const handleAcceptTask = async (taskID) => {
		await acceptTask(taskID, auth.userID, auth.token);
		setTasks(await getProjectTasks(project.id, auth.token));
	};
	const handleCompleteTask = async (taskID) => {
		await completeTask(taskID, auth.userID, auth.token);
		setTasks(await getProjectTasks(project.id, auth.token));
	};
	const taskActions = {
		setTaskForInfo: setTaskForInfo,
		setTaskForEdit: setTaskForEdit,
		deleteTask: handleDeleteTask,
		acceptTask: handleAcceptTask,
		completeTask: handleCompleteTask,
	};

	const sendTaskData = async (taskData) => {
		if (taskData.id.length > 0) {
			await updateTask(taskData.id, taskData, auth.token);
		} else {
			await addTask(project.id, taskData, auth.token);
		}

		setTasks(await getProjectTasks(project.id, auth.token));
	};

	return (
		<>
			<Box>
				{auth !== null && auth.role === "Manager" && (
					<Button
						variant="contained"
						className="py-3 mb-8 w-full"
						onClick={() => {
							setSelectedTask(null);
							setTaskElement("form");
							handleOpen();
						}}>
						{t("addTask")}
					</Button>
				)}

				{tasks !== null && (
					<Stack
						direction={"row"}
						divider={<Divider orientation="vertical" flexItem />}>
						<TaskColumn
							tasks={tasks}
							taskStatus="To do"
							taskActions={taskActions}
						/>
						<TaskColumn
							tasks={tasks}
							taskStatus="Doing"
							taskActions={taskActions}
						/>
						<TaskColumn
							tasks={tasks}
							taskStatus="Done"
							taskActions={taskActions}
						/>
					</Stack>
				)}
			</Box>
			<Dialog open={modalOpen} onClose={handleClose}>
				<DialogContent>
					{taskElement === "form" ? (
						<TaskForm
							task={selectedTask}
							sendTaskData={sendTaskData}
							handleClose={handleClose}
						/>
					) : (
						<TaskInfo
							task={selectedTask}
							handleClose={handleClose}
						/>
					)}
				</DialogContent>
			</Dialog>
		</>
	);
};

export default TasksTab;
