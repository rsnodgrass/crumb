# VHS Demo Tape for crumb
# https://github.com/charmbracelet/vhs

Output assets/demo.gif

Set Shell "bash"
Set FontSize 14
Set Width 900
Set Height 650
Set Theme "Catppuccin Mocha"
Set Padding 20

# Set clean prompt (white "% ")
Set TypingSpeed 50ms
Env PS1 "% "

# Initialize crumbs directory with README
Type "./crumb init"
Enter
Sleep 500ms

# Launch the TUI
Type "./crumb"
Enter
Sleep 1s

# Prompt field is focused first - start typing immediately
Type "Please refactor the authentication module to use JWT tokens instead of session cookies."
Enter
Enter
Type "Requirements:"
Enter
Type "- Use RS256 algorithm for signing"
Enter
Type "- Include refresh token rotation"
Enter
Type "- Add token blacklist for logout"
Sleep 500ms

# Tab to Output field (skip it)
Tab
Sleep 300ms

# Tab to Title field - add optional title
Tab
Sleep 300ms
Type "Refactor auth to JWT"
Sleep 400ms

# Tab to Tool selector (already defaults to Claude Code)
Tab
Sleep 300ms

# Tab to Tags field
Tab
Sleep 300ms

# Add tags using number keys (quick-select favorites)
Type "1"
Sleep 200ms
Type "3"
Sleep 200ms

# Also type a custom tag
Type "security"
Enter
Sleep 300ms

# Save with Ctrl+S
Ctrl+S
Sleep 2s

# Show the result
Sleep 500ms
