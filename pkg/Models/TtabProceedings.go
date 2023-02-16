package Models 

import (
)

type TtabProceedings struct {
    TtabProceedings string `xml:"ttab-proceedings"`
    Version struct {
        VersionNo     string `xml:"version-no"`
        VersionDate   string `xml:"version-date"`
    } `xml:"version"`
    ActionKeyCode       string `xml:"action-key-code"`
    TransactionDate     string `xml:"transaction-date"`
    ProceedingInfo      ProceedingInformation `xml:"proceeding-information"`
}

type ProceedingInformation struct {
   
    ProceedingEntry []struct {
        Number                 int       `xml:"number"`
        TypeCode               string    `xml:"type-code"`
        FilingDate             string    `xml:"filing-date"`
        EmployeeNumber         int       `xml:"employee-number"`
        InterlocutoryAttorney  string    `xml:"interlocutory-attorney-name"`
        LocationCode           string    `xml:"location-code"`
        DayInLocation          string    `xml:"day-in-location"`
        StatusUpdateDate       string    `xml:"status-update-date"`
        StatusCode             int       `xml:"status-code"`
        PartyInformation       struct {
            Party []struct {
                Identifier         int       `xml:"identifier"`
                RoleCode           string    `xml:"role-code"`
                Name               string    `xml:"name"`
                PropertyInformation struct {
                    Property []struct {
                        Identifier    int    `xml:"identifier"`
                        SerialNumber  int    `xml:"serial-number"`
                        MarkText      string `xml:"mark-text"`
                    } `xml:"property"`
                } `xml:"property-information"`
                AddressInformation struct {
                    ProceedingAddress []struct {
                        Identifier  int    `xml:"identifier"`
                        TypeCode    string `xml:"type-code"`
                        Name        string `xml:"name"`
                        Address1    string `xml:"address-1"`
                        City        string `xml:"city"`
                        State       string `xml:"state"`
                        Country     string `xml:"country"`
                        Postcode    string `xml:"postcode"`
                    } `xml:"proceeding-address"`
                } `xml:"address-information"`
            } `xml:"party"`
        } `xml:"party-information"`
        ProsecutionHistory struct {
            ProsecutionEntry []struct {
                Identifier   int    `xml:"identifier"`
                Code         int    `xml:"code"`
                TypeCode     string `xml:"type-code"`
                Date         string `xml:"date"`
                HistoryText  string `xml:"history-text"`
            } `xml:"prosecution-entry"`
        } `xml:"prosecution-history"`
    } `xml:"proceeding-entry"`
}