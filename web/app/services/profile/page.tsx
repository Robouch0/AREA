"use server";
import {getUserInfo, userInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/ui/profilePage";

export default async function Profile() {
     const data:userInfo = await getUserInfo();
    return <ProfilePage {...data}/>;
}
