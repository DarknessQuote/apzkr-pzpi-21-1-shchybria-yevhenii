import { IconButton, Menu, MenuItem } from "@mui/material";
import { useState } from "react";
import LanguageOutlinedIcon from "@mui/icons-material/LanguageOutlined";
import { useTranslation } from "react-i18next";

const languages = {
	en: {
		name: "English",
	},
	ua: {
		name: "Ukrainian",
	},
};

const ChangeLanguageMenu = () => {
	const [anchorEl, setAnchorEl] = useState(null);

	const { i18n } = useTranslation();

	return (
		<>
			<IconButton onClick={(e) => setAnchorEl(e.currentTarget)}>
				<LanguageOutlinedIcon className="text-white" />
			</IconButton>
			<Menu
				anchorEl={anchorEl}
				open={Boolean(anchorEl)}
				onClose={() => setAnchorEl(null)}>
				{Object.keys(languages).map((lang) => {
					return (
						<MenuItem
							key={lang}
							onClick={() => {
								i18n.changeLanguage(lang);
								setAnchorEl(null);
							}}>
							{languages[lang].name}
						</MenuItem>
					);
				})}
			</Menu>
		</>
	);
};

export default ChangeLanguageMenu;
