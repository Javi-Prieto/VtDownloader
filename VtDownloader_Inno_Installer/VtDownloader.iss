[Setup]
AppName=VtDownloader
AppVersion=0.1
DefaultDirName={commonpf}\VtDownloader
DefaultGroupName=VtDownloader
OutputBaseFilename=VtDownloaderInstaller
Compression=lzma
SolidCompression=yes

[Files]
Source: "myPath"; DestDir: "{app}\back\main\"; Flags: ignoreversion
Source: "myPath"; DestDir: "{app}\back\"; Flags: ignoreversion
Source: "myPath"; DestDir: "{app}\front\"; Flags: ignoreversion

[Icons]
Name: "{group}\VtDownloader"; Filename: "{app}\VtDownloader Setup 1.0.0.exe"
Name: "{userdesktop}\VtDownloader"; Filename: "{app}\VtDownloader Setup 1.0.0.exe"; Tasks: desktopicon

[Run]
Filename: "{app}\back\main\cmd.exe"; Description: "Run Backend"; Flags: runhidden shellexec
Filename: "{app}\front\VtDownloader Setup 1.0.0.exe"; Description: "Launch YourApp"; Flags: nowait postinstall skipifsilent

[Tasks]
Name: "desktopicon"; Description: "Create a desktop icon"; GroupDescription: "Additional icons:"
Name: "runbackend"; Description: "Run Backend on Startup"; GroupDescription: "Backend Options:"
