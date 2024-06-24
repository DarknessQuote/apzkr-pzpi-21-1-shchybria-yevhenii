import { useEffect, useState } from "react";
import {
	addCompany,
	deleteCompany,
	getCompanies,
	updateCompany,
} from "../services/companyService";
import {
	Box,
	Button,
	ButtonGroup,
	Dialog,
	DialogContent,
	Divider,
	List,
	ListItem,
	ListItemButton,
	ListItemText,
	Paper,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "../context/AuthContext";
import CompanyForm from "../components/CompanyForm";

const CompaniesPage = () => {
	const [companies, setCompanies] = useState([]);
	const [selectedCompany, setSelectedCompany] = useState(null);
	const [modalOpen, setModalOpen] = useState(false);

	const [auth] = useAuthContext();

	const { t } = useTranslation();

	useEffect(() => {
		const loadCompanies = async () => {
			const companiesList = await getCompanies();
			setCompanies(companiesList);
		};

		loadCompanies();
	}, []);

	const handleDeleteCompany = async (companyID) => {
		await deleteCompany(companyID, auth.token);
		setCompanies(await getCompanies());
	};

	const sendCompanyData = async (companyData) => {
		if (companyData.id.length > 0) {
			await updateCompany(companyData.id, companyData, auth.token);
		} else {
			await addCompany(companyData, auth.token);
		}

		setCompanies(await getCompanies());
	};

	const handleOpen = () => setModalOpen(true);
	const handleClose = () => setModalOpen(false);

	return (
		<>
			<Paper>
				<List>
					{companies.map((company, i) => {
						return (
							<Box key={company.id}>
								<ListItem className="flex justify-start">
									<ListItemText
										primary={company.name}
										secondary={company.owner}
										className="w-96 grow-0"
									/>
									<ListItemText
										secondary={company.email}
										className="grow"
									/>
									<ButtonGroup variant="contained">
										<ListItemButton
											onClick={() => {
												setSelectedCompany(company);
												handleOpen();
											}}>
											{t("edit")}
										</ListItemButton>
										<ListItemButton
											onClick={() =>
												handleDeleteCompany(company.id)
											}>
											{t("delete")}
										</ListItemButton>
									</ButtonGroup>
								</ListItem>
								{i < companies.length - 1 && <Divider />}
							</Box>
						);
					})}
				</List>
			</Paper>
			<Button
				variant="contained"
				className="py-3"
				onClick={() => {
					setSelectedCompany(null);
					handleOpen();
				}}>
				{t("addCompany")}
			</Button>
			<Dialog open={modalOpen} onClose={handleClose}>
				<DialogContent>
					<CompanyForm
						company={selectedCompany}
						sendCompanyData={sendCompanyData}
						handleClose={handleClose}
					/>
				</DialogContent>
			</Dialog>
		</>
	);
};

export default CompaniesPage;
