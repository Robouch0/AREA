"use server";
import {getUserInfo, getUserTokens, userInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/pages/profile/ProfilePage";

export default async function Profile() {
     const data:userInfo = await getUserInfo();
     data.providers = await getUserTokens();

    return <ProfilePage {...data} />;
}
