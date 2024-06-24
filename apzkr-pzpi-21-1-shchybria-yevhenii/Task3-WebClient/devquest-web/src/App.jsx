import { Route, Routes } from "react-router-dom";
import AuthPage from "./routes/AuthPage";
import HomePage from "./routes/HomePage";
import RootLayout from "./routes/RootLayout";
import CompaniesPage from "./routes/CompaniesPage";
import ProjectsPage from "./routes/ProjectsPage";
import ProjectPage from "./routes/ProjectPage";
import { useAuthContext } from "./context/AuthContext";
import RouteProtection from "./routes/RouteProtection";

const App = () => {
	const [auth] = useAuthContext();

	return (
		<Routes>
			<Route path="/">
				<Route element={<RootLayout />}>
					<Route index element={<HomePage />} />
					<Route path="auth" element={<AuthPage />} />
					<Route
						element={
							<RouteProtection
								authorizedRoles={["Admin"]}
								auth={auth}
							/>
						}>
						<Route path="companies" element={<CompaniesPage />} />
					</Route>
					<Route
						element={
							<RouteProtection
								authorizedRoles={["Manager", "Developer"]}
								auth={auth}
							/>
						}>
						<Route path="projects" element={<ProjectsPage />} />
						<Route path="project/:id" element={<ProjectPage />} />
					</Route>
				</Route>
			</Route>
		</Routes>
	);
};

export default App;
