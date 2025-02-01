# Single Responsibility Principle (SRP)

The Single Responsibility Principle (SRP) is one of the five SOLID principles of object-oriented design. It states that a class should have only one reason to change, meaning that a class should have only one job or responsibility. This principle helps to achieve a high level of cohesion within classes and makes the system easier to maintain and understand.

## Explanation

A class should only have one responsibility, and therefore only one reason to change. If a class has more than one responsibility, those responsibilities become coupled, and changes to one responsibility may affect or break the other responsibilities.

### Benefits of SRP

- **Improved Cohesion**: Classes are more focused and easier to understand.
- **Easier Maintenance**: Changes in requirements affect fewer classes.
- **Enhanced Reusability**: Classes with a single responsibility are more likely to be reusable in different contexts.

## Example

Consider a class that handles both user authentication and user data management. According to SRP, these responsibilities should be separated into different classes.

### Before SRP

```
+----------------------+
| UserManager          |
+----------------------+
| - AuthenticateUser() |
| - GetUserData()      |
| - SaveUserData()     |
+----------------------+
```

### After SRP

```
+----------------------+
| AuthManager          |
+----------------------+
| - AuthenticateUser() |
+----------------------+

+----------------------+
| UserDataManager      |
+----------------------+
| - GetUserData()      |
| - SaveUserData()     |
+----------------------+
```

By separating the responsibilities, each class now has a single responsibility, making the system more modular and easier to maintain.
