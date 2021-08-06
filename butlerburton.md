% butler-burton(1) butler-burton
% Patrik Olin <patrik@olin.work>
% August 2021

# NAME

butler-burton - a smartish utility to manage your BL time sheet

# SYNOPSIS

**butler-burton** [global options] _command_ [command options] [arguments...]

# DESCRIPTION

**butler-burton** helps you manage and fill out your time sheet by having you
check in and check out at appropriate times. Running **butler-burton** without any
command line parameters will display its version and help section.

# GLOBAL OPTIONS

**-h**, **--help**
: Displays the super friendly help section.

**-v**, **--version**
: Displays the software version.

**-s**, **--silent**
: Prevents **butler-burton** from sending Teams messages for this command.

# EXAMPLES

**butler-burton**
: Displays super friendly help section and version information.

**butler-burton ci | butler-burton check-in**
: Standard check-in procedure. Will send a Teams message saying you are checked in and write your check-in time to the time sheet.

**butler-burton co | butler-burton check-out**
: Standard check-out procedure. Will send a Teams message saying you are check out and write your check-out time to the time sheet.

**butler-burton ts d && butler-burton ci && butler-burton ts u**
: Downloads your time sheet from Sharepoint, checks you in and updates the time sheet with check-in time before uploading the time sheet to Sharepoint again.

**butler-burton ts d && butler-burton co -c -o && butler-burton ts u**
: Downloads your time sheet from Sharepoint, checks you out and updates the time sheet with your check-out time, catered lunch and writes extra hours to the overtime column and then uploads your time sheet to Sharepoint.

# BUGS

I should think so.

# COPYRIGHT

No, do with it as your please and use it at your own risk.
