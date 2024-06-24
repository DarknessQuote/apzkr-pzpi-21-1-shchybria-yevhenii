export const getUser = async (userID) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/user/${userID}`,
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

export const getRole = async (roleID) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/role/${roleID}`,
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
