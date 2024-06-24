import { createContext, useContext, useEffect, useState } from "react";
import { refresh } from "../services/authService";

const AuthContext = createContext();

export const useAuthContext = () => {
	return useContext(AuthContext);
};

export const AuthProvider = ({ children }) => {
	const [auth, setAuth] = useState(null);

	useEffect(() => {
		const refreshAuthData = async () => {
			try {
				const refreshData = await refresh();

				if (refreshData == null) {
					return;
				}

				setAuth({
					token: refreshData.tokens.access_token,
					userID: refreshData.user_id,
					role: refreshData.role,
				});
			} catch (err) {
				console.error(err);
			}
		};

		refreshAuthData();
	}, []);

	return (
		<AuthContext.Provider value={[auth, setAuth]}>
			{children}
		</AuthContext.Provider>
	);
};
