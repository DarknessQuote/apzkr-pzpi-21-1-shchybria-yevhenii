export const getCompanies = async () => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/companies`,
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

export const getCompany = async (companyID) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/companies/${companyID}`,
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

export const addCompany = async (addCompanyData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: addCompanyData.name,
			owner: addCompanyData.owner,
			email: addCompanyData.email,
		};

		const reqOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/companies`,
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

export const updateCompany = async (companyID, updateCompanyData, token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/json");
		headers.append("Authorization", `Bearer ${token}`);

		const reqBody = {
			name: updateCompanyData.name,
			owner: updateCompanyData.owner,
			email: updateCompanyData.email,
		};

		const reqOptions = {
			method: "PUT",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(reqBody),
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/companies/${companyID}`,
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

export const deleteCompany = async (companyID, token) => {
	try {
		const headers = new Headers();
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "DELETE",
			credentials: "include",
			headers: headers,
		};

		const responseJSON = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/companies/${companyID}`,
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
