export const register = async (registerData) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqBody = {
			username: registerData.username,
			first_name: registerData.firstName,
			last_name: registerData.lastName,
			password: registerData.password,
			role_id: registerData.roleID,
			company_id: registerData.companyID,
		};

		const reqOptions = {
			method: "POST",
			body: JSON.stringify(reqBody),
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/auth/register`,
			reqOptions
		);
		const response = await responseJSON.json();

		if (response.error) {
			throw new Error(response.message);
		}

		return response.data;
	} catch (err) {
		throw err;
	}
};

export const login = async (loginData) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqBody = {
			username: loginData.username,
			password: loginData.password,
		};

		const reqOptions = {
			method: "POST",
			body: JSON.stringify(reqBody),
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/auth/login`,
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

export const refresh = async () => {
	try {
		const reqOptions = {
			method: "POST",
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/auth/refresh`,
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

export const logout = async () => {
	try {
		const reqOptions = {
			method: "DELETE",
			credentials: "include",
		};

		await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/auth/logout`,
			reqOptions
		);
	} catch (err) {
		throw err;
	}
};

export const getRoles = async () => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/auth/roles`,
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
