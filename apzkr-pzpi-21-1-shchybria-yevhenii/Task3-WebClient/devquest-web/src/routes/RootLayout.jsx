import { Box, Container } from "@mui/material";
import Header from "../components/Header";
import { Outlet } from "react-router-dom";

const RootLayout = () => {
	return (
		<Box>
			<Header />
			<Container maxWidth="lg" className="py-5 flex flex-col gap-5">
				<Outlet />
			</Container>
		</Box>
	);
};

export default RootLayout;
