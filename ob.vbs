Set objWMI = GetObject("winmgmts:\\.\root\cimv2")
Set colOS = objWMI.ExecQuery("SELECT * FROM Win32_OperatingSystem")

For Each objOS in colOS
    If CInt(objOS.BuildNumber) < 22000 Then
        MsgBox "Oops... This app can be runned only on Windows 11. Upgrade your system and try again.", vbCritical, "Device Manager"
    End If
Next
