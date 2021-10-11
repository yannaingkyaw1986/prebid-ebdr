package engagebdr

import (
    "testing"
    "text/template"

    "github.com/prebid/prebid-server/privacy"
    "github.com/prebid/prebid-server/privacy/gdpr"
    "github.com/stretchr/testify/assert"
)

func TestEngageBDRSyncer(t *testing.T) {
    syncURL := "https://match.bnmla.com/usersync?sspid=9000&redir=localhost%2Fsetuid%3Fbidder%3Dengagedr%26gdpr%3D%5BGDPR%5D%26gdpr_consent%3D%5BGDPRConsent%5D%26uid%3D%5Buid%5D"
    syncURLTemplate := template.Must(
        template.New("sync-template").Parse(syncURL),
    )   

    syncer := NewEngageBDRSyncer(syncURLTemplate)
    syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
        GDPR: gdpr.Policy{
            Signal: "0",
        },  
    })  

    assert.NoError(t, err)
    assert.Equal(t, "//match.bnmla.com/usersync?sspid=9000&redir=localhost%2Fsetuid%3Fbidder%3Dengagebdr%26gdpr%3D%26gdpr_consent%3D%26uid%3D%5BUID%5D", syncInfo.URL)
    assert.Equal(t, "redirect", syncInfo.Type)
    assert.EqualValues(t, 0, syncer.GDPRVendorID())
    assert.Equal(t, false, syncInfo.SupportCORS)
}
