1. Structure: In index.html, replace the username label with "Email:" and change the input type from "text" to "email", updating its id and name attributes to "email"
2. Structure: Remove the entire `<div id="registration-fields">` section from index.html along with all its inner fields (firstname, lastname, email, dateofbirth, confirmpassword) since registration will only require email and password
3. Logic: In script.js, update the toggle button event listener to remove any code that shows/hides the `#registration-fields` container and toggles the `.register-mode` class on the form container; keep only the form title and button text changes
4. Logic: In script.js, refactor the form submit handler: replace the `username` variable with `email` from the new email input, remove all registration-specific field extraction and validation (firstname, lastname, email inside registration, dateofbirth, confirmpassword, password mismatch), and for registration mode simply extract and log/process the `email` and `password` fields
5. Styling: Clean up styles.css by removing the entire `.form-row` rule set, the `#registration-fields` rule set, the `.form-container.register-mode` class, and the associated `.form-row` media query since these elements no longer exist
6. Polish: Update any remaining console.log statements or comments in script.js to reference "email" instead of "username" for consistency

###Here is the step by step plan

# plan.md

## High-Level Objectives

1. Structure: In index.html, replace the username label with "Email:" and change the input type from "text" to "email", updating its id and name attributes to "email"
2. Structure: Remove the entire `<div id="registration-fields">` section from index.html along with all its inner fields (firstname, lastname, email, dateofbirth, confirmpassword) since registration will only require email and password
3. Logic: In script.js, update the toggle button event listener to remove any code that shows/hides the `#registration-fields` container and toggles the `.register-mode` class on the form container; keep only the form title and button text changes
4. Logic: In script.js, refactor the form submit handler: replace the `username` variable with `email` from the new email input, remove all registration-specific field extraction and validation (firstname, lastname, email inside registration, dateofbirth, confirmpassword, password mismatch), and for registration mode simply extract and log/process the `email` and `password` fields
5. Styling: Clean up styles.css by removing the entire `.form-row` rule set, the `#registration-fields` rule set, the `.form-container.register-mode` class, and the associated `.form-row` media query since these elements no longer exist
6. Polish: Update any remaining console.log statements or comments in script.js to reference "email" instead of "username" for consistency

## Detailed Implementation Steps

1. Open `index.html` and change the login field label and input to use email instead of username:
   - Locate the line `<label for="username">Username:</label>` and replace it with `<label for="email">Email:</label>`
   - Locate the line `<input type="text" id="username" name="username" required>` and replace it with `<input type="email" id="email" name="email" required>`

2. Open `index.html` and delete the entire registration fields section:
   - Remove everything from the opening `<div id="registration-fields" style="display: none;">` through its closing `</div>` (the line immediately before the `<!-- Buttons -->` comment). All inner `form-row` divs (firstname, lastname, email, dateofbirth, confirmpassword) must be removed.

3. Open `script.js` and update the toggle button event listener to remove registration field toggling and container class manipulation:
   - In the `document.getElementById('toggle-btn').addEventListener('click', function() { ... })` block, delete the following lines:
     - `const registrationFields = document.getElementById('registration-fields');`
     - `const formContainer = document.querySelector('.form-container');`
   - Inside the `if (isLoginMode)` block, delete:
     - `registrationFields.style.display = 'none';`
     - `formContainer.classList.remove('register-mode');`
   - Inside the `else` block, delete:
     - `registrationFields.style.display = 'block';`
     - `formContainer.classList.add('register-mode');`
   - The resulting callback should only toggle `isLoginMode`, update `formTitle.textContent` to 'Login' or 'Register', and update `toggleBtn.textContent` to 'Switch to Register' or 'Switch to Login'.

4. Open `script.js` and refactor the form submit handler:
   - Replace `const username = document.getElementById('username').value;` with `const email = document.getElementById('email').value;`
   - Delete the entire `if (!isLoginMode) { ... }` block that starts after `setTimeout(() => {` and before `console.log('Login Data:', ...)`. This includes the lines that fetch `firstname`, `lastname`, the inner `email`, `dateofbirth`, `confirmpassword`, the validation checks, and the `console.log('Registration Data:', ...)` inside that block.
   - Replace the remaining login/registration logging with:
     ```javascript
     if (isLoginMode) {
       console.log('Login Data:', { email, password });
     } else {
       console.log('Registration Data:', { email, password });
     }
     ```
   - The loading button state (`setTimeout` wrapper, original text storage, disabled logic) and the reset after 2 seconds must remain intact.

5. Open `styles.css` and delete the following CSS rules entirely:
   - The entire `.form-row { ... }` rule set and all its nested declarations (`.form-row label`, `.form-row input`)
   - The `#registration-fields { ... }` rule set
   - The `.form-container.register-mode { ... }` rule set
   - The `@media (max-width: 600px) { ... }` block that targets `.form-row` (the whole media query)

6. Open `script.js` and ensure no stale 'username' references remain:
   - Verify that inside the submit handler's `console.log` statements (already updated in step 4) use `email` instead of `username`.
   - Check for any other comments or log statements referencing `username` and replace them with `email` if found (the only occurrences were in the submit handler).