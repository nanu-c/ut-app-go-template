import QtQuick 2.4
import Ubuntu.Components 1.3

MainView {

    // Note! applicationName needs to match the "name" field of the click manifest
    applicationName: "gotestapp.nanuc"
    // If you set a objectName it can be found from go
    objectName: "mainView"


    automaticOrientation: false


    width: units.gu(100)
    height: units.gu(75)


    Page {
        title: i18n.tr("Go Testapp")

        Column {
            id: col
            spacing: units.gu(1)
            anchors {
                margins: units.gu(2)
                fill: parent
            }
            Label {
              id: outLabel
              objectName: "label"
              text: testvar.message // to export go functions they have to be
                                    // in uppercase but in qml the first letter
                                    // of the function has to be lower case.
                                    // the qml<->go bridge takes care of that
            }
            Text {
              id: output
              width: parent.width
              height: parent.height / 1.5
              text: testvar.output
              wrapMode: Text.WordWrap
           }
            Button {
                id: refreshBtn
                onClicked: testvar.getMessage()
                text: "Refresh"
                color: "green"
            }
      }
    }
}
