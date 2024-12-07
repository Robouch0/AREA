"use server";
import {getUserInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/ui/profilePage";

export default async function Profile() {
     const data = await getUserInfo();
    return <ProfilePage userData={data}/>;
}
