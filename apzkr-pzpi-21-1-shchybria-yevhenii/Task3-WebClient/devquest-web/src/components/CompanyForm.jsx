import { Box, Button, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

const CompanyForm = ({ company, sendCompanyData, handleClose }) => {
	const [id, setID] = useState("");
	const [name, setName] = useState("");
	const [owner, setOwner] = useState("");
	const [email, setEmail] = useState("");

	const { t } = useTranslation();

	useEffect(() => {
		if (company !== null) {
			setID(company.id);
			setName(company.name);
			setOwner(company.owner);
			setEmail(company.email);
		}
	}, [company]);

	const handleSubmit = async (e) => {
		e.preventDefault();

		const companyData = {
			id: id,
			name: name,
			owner: owner,
			email: email,
		};

		await sendCompanyData(companyData);
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
						name="owner"
						label={t("owner")}
						variant="outlined"
						value={owner}
						onChange={(e) => setOwner(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
					<TextField
						required
						name="email"
						label={t("email")}
						variant="outlined"
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						InputLabelProps={{ shrink: true }}
					/>
				</Box>
				<Button variant="contained" type="submit">
					{t("save")}
				</Button>
			</Box>
		</form>
	);
};

export default CompanyForm;
