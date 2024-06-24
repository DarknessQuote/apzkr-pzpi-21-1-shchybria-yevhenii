import { getUser } from "./userService";

export const getProject = async (projectID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/projects/${projectID}`,
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

export const getManagerProjects = async (managerID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/projects/manager/${managerID}`,
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

export const getDeveloperProjects = async (developerID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/projects/developer/${developerID}`,
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

export const getProjectDevelopers = async (projectID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/projects/developers/${projectID}`,
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

export const getAvailableDevelopers = async (managerID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/projects/available-developers/${managerID}`,
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

export const addProject = async (managerID, addProjectData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const manager = await getUser(managerID);

		const reqBody = {
			name: addProjectData.name,
			description: addProjectData.description,
			company_id: manager.company_id,
			manager_id: manager.id,
		};

		const reqOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/projects`,
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

export const updateProject = async (projectID, updateProjectData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: updateProjectData.name,
			description: updateProjectData.description,
		};

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/projects/${projectID}`,
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

export const deleteProject = async (projectID, token) => {
	const headers = new Headers();
	headers.append("Authorization", `Bearer ${token}`);

	try {
		const reqOptions = {
			method: "DELETE",
			credentials: "include",
			headers: headers,
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/projects/${projectID}`,
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

export const addDeveloperToProject = async (projectID, developerID, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/projects/developers?projectID=${projectID}&developerID=${developerID}`,
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

export const removeDeveloperFromProject = async (
	projectID,
	developerID,
	token
) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "DELETE",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/projects/developers?projectID=${projectID}&developerID=${developerID}`,
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
