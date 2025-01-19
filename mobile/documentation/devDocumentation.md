# AREA Mobile Project Documentation

## Project Structure

```
.
├── api/
├── core/
├── main.dart
├── pages/
├── services/
└── widgets/
```

## Directory Overview

### 📁 api/types/
Contains data models for various API requests:
- Area management (creation, activation)
- Authentication
- User profile
- OAuth providers

### 📁 core/router/
Application routing management:
- `app_router.dart`: Route configuration
- `route_names.dart`: Route name constants

### 📁 pages/
Main application screens:
- Authentication (Login/Register)
- Home
- Profile
- Areas (creation and listing)
- Server configuration

### 📁 services/
Business logic services:
- `api/`: API call services (areas, auth, profile, server)
- `storage/`: Local storage management (auth)

### 📁 widgets/
Reusable UI components:
- Authentication components (buttons, fields)
- Navigation (bottom navbar)
- Layout (main scaffold)
- OAuth (connection buttons, webview)

## Key Information for Developers

### Service Architecture
- API-oriented service architecture
- Clear separation between business logic (services) and user interface (pages/widgets)
- Local storage implementation for auth data persistence

### Authentication Management
- Support for traditional authentication (email/password)
- OAuth integration with various providers
- Secure token storage

### Development Guidelines
- Use custom widgets from the `widgets/` directory to maintain UI consistency
- Follow existing structure when adding new pages or features
- Respect separation of concerns:
    - Pages for views
    - Services for business logic
    - Models for data structures

### Extension Points
To add a new feature:
1. Create necessary data models in `api/types/`
2. Implement corresponding API services in `services/api/`
3. Create the page in the `pages/` directory
4. Add the route in `core/router/`

### Best Practices
1. **Code Organization**
    - Keep files focused and single-purpose
    - Use clear, descriptive naming conventions
    - Group related functionality together

2. **State Management**
    - Handle state consistently throughout the app
    - Use appropriate state management solutions for different scenarios
    - Keep state logic separate from UI logic

3. **API Integration**
    - Always use the service layer for API calls
    - Handle errors gracefully
    - Implement proper loading states
