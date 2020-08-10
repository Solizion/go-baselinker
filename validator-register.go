package baselinker

func (baseLinker *BaseLinker) registerValidationMethods() {
	baseLinker.validator.RegisterValidation("is-journal-log-types", validateJournalTypes)
}
