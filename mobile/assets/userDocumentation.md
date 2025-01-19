# User Documentation Walkthrough

## Introduction
- **Purpose of the Documentation**: This documentation provides a comprehensive guide to using the appliation, detailing its features and functionalities.
- **Overview of Features**: The application includes multiple pages: My Area, Profile, Create, Login, and Register.

## 1. My Area Page
### Description
The My Area page allows users to view and manage their created Area Cards.

On each of your Area Card you can see the action and first reaction description.

Aswell as the Category of service his action is coming from (ex: Github)
# Functionalities
## **Enable/Disable Areas**:

- You can click the switch on the area card to disable or enable them.

![How to enable or disable areas](my_area.png)


## Visual overview of the My Areas page

![Full my area page](my_area_page.png)

---

## 2. Profile Page
### Description

The Profile page enables users to manage their personal information and link external accounts.

### Sections
- **Update User Data**:

- ![How to edit the profile informations](profile_edition.png)

    - The top side of the profile page allows user to edit their personnal informations with input fields.

    - Write down the new value in the corresponding field, then press the 'Save' button to update your datas on Area servers.


- **Link OAuth Accounts**:

![How to link oauth from profile tab](profile_oauth.png)

- The bottom side of the profile page allows you to scroll through services (e.g., GitHub, Google, Spotify).
- You can click on the Unlink button to make Area forgot about your linked account on the designated service.
- Or press the Link button to connect an external account with Oauth validation.

---

## 3. Create Page
### Description

The Create page allows users to set up automated actions by choosing services and using microservices from them.

On this page you can choose one actions and add multiple reactions to it.

### Functionality
- **Choose Service for Action/Reaction**:

- Firstly, you need to select a service on the dropdown.

- **Select Microservices**:

- After choosing a service, microservices card with description will be displayed. You can click on them to select one.

![Create an action from create tab](create_action.png)

- **Fill out the values**:

    - For the area to be created the required fields of the microservices need to be filled.
    - Some fields need texts in it
    - In some that needs time data you can select it in the calendar directly
    - If you need more information about an ingredient you can read the brief description of what we are waiting for as the hint text.
    - If there is any variable that you can use on the next reaction, simply provide the variable like this {{.variableName}} inside text inputs.

![How to fill ingredients values](action_ingredients.png)

- After you filled out everything you can click on the create AREA button to finish the process and create your process automation !

![Create button from create tab](create_area_button.png)

- **Add chained reactions**:

In our application you can create area with one action and then add multiple chained reactions to it.
For this just click on the [**+ Add Reaction**] button
then fill out the new inputfield needed before creating your area.

![Add a new chained reaction](add_reaction_area.png)

### Visual overview of the Create page
![Full create page](create_page.png)
---

## 4. Login Page
### Description
The Login page is where users access their accounts.

### Steps to Log In
- To log-in you can either enter your credentials (email and password).
- Or you can connect with Oauth (Discord, GitHub)
- If you do not own an Area account for the moment, at the bottom of the page you cand find a link to go to the register page, to create an account.

![How to login on the app](login_page.png)

---

## 5. Register Page
### Description
The Register page allows new users to create an account.

### Steps to Register
- Filling out registration form fields (firstname, lastname, email, password).
- If you want to do a quick sign-up you can click on one of the three bottom buttons to connect through an Oauth of one of your already existing account (Github, Discord)

We will then take your email on this account as a data on your profile.

![How to register on the app](register_page.png)

## Conclusion

This documentation act as a user guide / walkthrough to help you naviguate on our application.
