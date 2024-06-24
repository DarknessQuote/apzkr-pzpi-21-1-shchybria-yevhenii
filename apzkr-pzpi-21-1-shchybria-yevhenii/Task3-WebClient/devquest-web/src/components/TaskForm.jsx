import {
	Box,
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	TextField,
} from "@mui/material";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { getTaskCategories } from "../services/taskService";
import { useAuthContext } from "../context/AuthContext";
import { DateTimePicker } from "@mui/x-date-pickers";
import dayjs from "dayjs";

const TaskForm = ({ task, sendTaskData, handleClose }) => {
	const [categories, setCategories] = useState([]);
	const [id, setID] = useState("");
	const [name, setName] = useState("");
	const [description, setDescription] = useState("");
	const [points, setPoints] = useState("");
	const [expectedTime, setExpectedTime] = useState(dayjs());
	const [categoryID, setCategoryID] = useState("");

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		if (task !== null) {
			setID(task.id);
			setName(task.name);
			setDescription(task.description);
			setPoints(task.points);
			setExpectedTime(dayjs(task.expected_time));
			setCategoryID(task.category_id);
		}

		const fetchTaskCategories = async () => {
			if (auth !== null) {
				const categories = await getTaskCategories(auth.token);
				setCategories(categories);
			}
		};
		fetchTaskCategories();
	}, [task, auth]);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const taskData = {
			id: id,
			name: name,
			description: description,
			points: Number(points),
			expectedTime: expectedTime.toISOString(),
			categoryID: categoryID,
		};

		await sendTaskData(taskData);
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
						name="points"
						type="number"
						label={t("points")}
						variant="outlined"
						value={points}
						onChange={(e) => setPoints(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
					<DateTimePicker
						label={t("expectedTime")}
						name="expectedTime"
						value={expectedTime}
						onChange={(newTime) => setExpectedTime(newTime)}
					/>
					<FormControl>
						<InputLabel>{t("category")}</InputLabel>
						<Select
							required
							title="role"
							label={t("role")}
							value={categoryID}
							onChange={(e) => setCategoryID(e.target.value)}>
							<MenuItem selected value="">
								{t("categorySelect")}
							</MenuItem>
							{categories.map((category) => {
								return (
									<MenuItem
										value={category.id}
										key={category.id}>
										{t(`${category.name}`)}
									</MenuItem>
								);
							})}
						</Select>
					</FormControl>
				</Box>
				<Button variant="contained" type="submit">
					{t("save")}
				</Button>
			</Box>
		</form>
	);
};

export default TaskForm;
