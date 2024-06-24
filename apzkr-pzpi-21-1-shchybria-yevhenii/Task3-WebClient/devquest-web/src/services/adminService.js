export const dataBackup = async (token) => {
	try {
		const headers = new Headers();
		headers.append("Content-Type", "application/x-tar");
		headers.append("Authorization", `Bearer ${token}`);

		const reqOptions = {
			method: "GET",
			headers: headers,
			credentials: "include",
		};

		const response = await fetch(
			`${process.env.REACT_APP_BACKEND_URL}/admin/data-backup`,
			reqOptions
		);

		const contentType = response.headers.get("Content-Type");
		if (contentType && contentType.indexOf("application/json") !== -1) {
			const responseJson = await response.json();
			throw new Error(responseJson.message);
		}

		const responseBlob = await response.blob();

		const fileurl = window.URL.createObjectURL(new Blob([responseBlob]));
		const link = document.createElement("a");
		link.href = fileurl;
		link.setAttribute("download", `devquest-data-backup.tar`);

		document.body.appendChild(link);
		link.click();
		link.parentNode.removeChild(link);
	} catch (err) {
		throw err;
	}
};
