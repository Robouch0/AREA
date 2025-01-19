import {
    FaGithub,
    FaGoogle,
    FaTwitter,
    FaDiscord,
    FaInstagram,
    FaYoutube,
    FaMailBulk,
    FaBrain,
    FaClock, FaSpotify, FaTrello, FaChessBoard, FaGitlab, FaCloud, FaCoins
} from "react-icons/fa";
import {IconType} from "react-icons";

export const serviceIcons: { [key: string]: IconType } = {
    github: FaGithub,
    google: FaGoogle,
    gitlab: FaGitlab,
    twitter: FaTwitter,
    discord: FaDiscord,
    instagram: FaInstagram,
    youtube: FaYoutube,
    tweeter: FaTwitter,
    outlook: FaMailBulk,
    spotify: FaSpotify,
    hf: FaBrain,
    dt: FaClock,
    asana: FaTrello,
    miro: FaChessBoard,
    weather: FaCloud,
    crypto: FaCoins,
};
