export const getProjectTasks = async (projectID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/${projectID}`,
			reqOptions
		);

		const response = await responseJSON.json();

		if (response === null) {
			return [];
		}

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const addTask = async (projectID, addTaskData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: addTaskData.name,
			description: addTaskData.description,
			points: addTaskData.points,
			expected_time: addTaskData.expectedTime,
		};

		const reqOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/${projectID}?categoryID=${addTaskData.categoryID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const updateTask = async (taskID, updateTaskData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: updateTaskData.name,
			description: updateTaskData.description,
			points: updateTaskData.points,
			expected_time: updateTaskData.expectedTime,
		};

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskID}?categoryID=${updateTaskData.categoryID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const deleteTask = async (taskID, token) => {
	try {
		const headers = new Headers();
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "DELETE",
			credentials: "include",
			headers: headers,
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const acceptTask = async (taskID, developerID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/accept?taskID=${taskID}&developerID=${developerID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const completeTask = async (taskID, developerID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/complete?taskID=${taskID}&developerID=${developerID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const getTaskCategories = async (token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/categories`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response === null) {
			return [];
		}

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const getTaskCategory = async (categoryID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/categories/${categoryID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response === null) {
			return null;
		}

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};

export const getTaskStatus = async (statusID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/tasks/status/${statusID}`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response === null) {
			return null;
		}

		if (response.error) {
			throw new Error(response.message);
		}

		return response;
	} catch (err) {
		throw err;
	}
};
