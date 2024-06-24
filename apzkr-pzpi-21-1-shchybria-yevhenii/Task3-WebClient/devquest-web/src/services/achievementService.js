export const getProjectAchievements = async (projectID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/achievements/project/${projectID}`,
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

export const getDeveloperAchievements = async (developerID, token) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/achievements/developer/${developerID}`,
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

export const addAchievement = async (projectID, addAchievementData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: addAchievementData.name,
			description: addAchievementData.description,
			points: addAchievementData.points,
		};

		const reqOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/achievements/${projectID}`,
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

export const updateAchievement = async (
	achievementID,
	updateAchievementData,
	token
) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: updateAchievementData.name,
			description: updateAchievementData.description,
			points: updateAchievementData.points,
		};

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/achievements/${achievementID}`,
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

export const deleteAchievement = async (achievementID, token) => {
	try {
		const headers = new Headers();
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "DELETE",
			credentials: "include",
			headers: headers,
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/achievements/${achievementID}`,
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

export const giveAchievement = async (
	projectID,
	achievementID,
	developerID,
	token
) => {
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
			`${process.env.REACT_APP_BACKEND_URL}/achievements/give?projectID=${projectID}&achievementID=${achievementID}&developerID=${developerID}`,
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
