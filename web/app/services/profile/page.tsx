"use server";
import {getUserInfo, userInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/pages/profile/ProfilePage";

export default async function Profile() {
     const data:userInfo = await getUserInfo();
    return <ProfilePage {...data}/>;
}
